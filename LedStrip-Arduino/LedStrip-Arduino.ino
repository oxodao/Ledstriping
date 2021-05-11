#define REDUCED_MODES // sketch too big for Arduino Leonardo flash, so invoke reduced modes
#include <WS2812FX.h>
#include <ArduinoJson.h>

#define LED_COUNT 15
#define LED_PIN 7
#define MAX_NUM_CHARS 15 // maximum number of characters read from the Serial Monitor

// Parameter 1 = number of pixels in strip
// Parameter 2 = Arduino pin number (most are valid)
// Parameter 3 = pixel type flags, add together as needed:
//   NEO_KHZ800  800 KHz bitstream (most NeoPixel products w/WS2812 LEDs)
//   NEO_KHZ400  400 KHz (classic 'v1' (not v2) FLORA pixels, WS2811 drivers)
//   NEO_GRB     Pixels are wired for GRB bitstream (most NeoPixel products)
//   NEO_RGB     Pixels are wired for RGB bitstream (v1 FLORA pixels, not v2)
//   NEO_RGBW    Pixels are wired for RGBW bitstream (NeoPixel RGBW products)
WS2812FX ws2812fx = WS2812FX(LED_COUNT, LED_PIN, NEO_GRB + NEO_KHZ400);

void(* reboot) (void) = 0;

struct State {
  uint32_t Color;
  uint8_t Brightness;
  uint8_t Mode;
  uint16_t Speed;

  bool Sparking;
  unsigned long SparkStart;
  unsigned long SparkDuration;

  int8_t Fading;
  uint8_t OriginalBrightness;
};

State currentState;
State previousState;

State getCurrentState() {
  return State {
    ws2812fx.getColor(),
    ws2812fx.getBrightness(),
    ws2812fx.getMode(),
    ws2812fx.getSpeed(),
    currentState.Sparking,
    currentState.SparkStart,
    currentState.SparkDuration,
    currentState.Fading,
    currentState.OriginalBrightness,
  };
}

void printState(State s) {
  StaticJsonDocument<200> doc;
  doc["color"] = s.Color;
  doc["brightness"] = s.Brightness;
  doc["mode"] = s.Mode;
  doc["speed"] = s.Speed;

  doc["sparking"] = s.Sparking;
  doc["spark_start"] = s.SparkStart;
  doc["spark_duration"] = s.SparkDuration;

  doc["fading"] = s.Fading;
  doc["original_brightness"] = s.OriginalBrightness;
  
  serializeJson(doc, Serial);
  Serial.println();
}

void setState(State s) {
  ws2812fx.setColor(s.Color);
  ws2812fx.setBrightness(s.Brightness);
  ws2812fx.setMode(s.Mode);
  ws2812fx.setSpeed(s.Speed);
}

void fadeIn() {
  unsigned long leftTime = millis() - currentState.SparkStart;
  float percentage = (float)leftTime / currentState.SparkDuration;

  uint8_t val = 0;
  if (currentState.OriginalBrightness * percentage > currentState.OriginalBrightness) {
      val = currentState.OriginalBrightness;
  } else {
    val = currentState.OriginalBrightness * percentage;
  }

  if (val == currentState.OriginalBrightness) {
    currentState.Fading = false;
  }

  currentState.Brightness = val;
}

void fadeOut() {
  unsigned long leftTime = millis() - currentState.SparkStart;
  float percentage = 1.0 - ((float)leftTime / currentState.SparkDuration);

  uint8_t val = 0;
  if (currentState.OriginalBrightness * percentage < 0) {
      val = 0;
  } else {
    val = currentState.OriginalBrightness * percentage;
  }

  if (val == 0) {
    currentState.Fading = false;
  }

  currentState.Brightness = val;
}

char cmd[MAX_NUM_CHARS];    // char[] to store incoming serial commands
bool cmd_complete = false;  // whether the command string is complete

void setup() {
  Serial.begin(9600);
  
  currentState = State {
    0x007BFF,
    30,
    FX_MODE_STATIC,
    1000,
    false,
    0,
    1000
  };

  currentState.OriginalBrightness = currentState.Brightness;
  
  previousState = currentState;

  ws2812fx.init();
  setState(currentState);
  ws2812fx.start();
}


void loop() {
  ws2812fx.service();

  recvChar();

  if(cmd_complete) {
    process_command();
  }

  if (currentState.Sparking && (currentState.SparkStart + currentState.SparkDuration) <= millis()) {
    currentState = previousState;
    setState(previousState);
  }

  if (currentState.Fading != 0) {
    if (currentState.Fading == 1) {
      fadeIn();
    } else {
      fadeOut();
    }
    setState(currentState);
  }
}

/*
 * Checks received command and calls corresponding functions.
 */
void process_command() {
  bool hasFullAnswer = false;
  bool isoke = false;

  //#region Brightness
  if (strcmp(cmd, "b") == 0) {
    Serial.println(ws2812fx.getBrightness());

    isoke = true;
    hasFullAnswer = true;
  }

  if (strncmp(cmd, "b ", 2) == 0) {
    currentState.Brightness = (uint8_t)atoi(cmd + 2);
    currentState.OriginalBrightness = currentState.Brightness;
    isoke = true;
  }
  //#endregion

  //#region Fade
  if (strcmp(cmd, "fadein") == 0) {
    currentState.SparkStart = millis();
    currentState.Fading = 1;

    isoke = true;
  }

  if (strcmp(cmd, "fadeout") == 0) {
    currentState.SparkStart = millis();
    currentState.Fading = -1;

    isoke = true;
  }
  //#endregion

  //#region Speed
  if (strcmp(cmd, "s") == 0) {
    Serial.println(ws2812fx.getSpeed());

    isoke = true;
    hasFullAnswer = true;
  }

  if (strncmp(cmd,"s ",2) == 0) {
    currentState.Speed = (uint16_t)atoi(cmd + 2);
    isoke = true;
  }
  //#endregion

  //#region Duration
  if (strcmp(cmd, "d") == 0) {
    Serial.println(currentState.SparkDuration);

    isoke = true;
    hasFullAnswer = true;
  }

  if (strncmp(cmd,"d ",2) == 0) {
    currentState.SparkDuration = (uint16_t)atoi(cmd + 2);
    isoke = true;
  }
  //#endregion

  //#region Mode
  if (strcmp(cmd, "m") == 0) {
    Serial.println(ws2812fx.getMode());

    isoke = true;
    hasFullAnswer = true;
  }

  if (strncmp(cmd,"m ",2) == 0) {
    currentState.Mode = (uint8_t)atoi(cmd + 2);
    isoke = true;
  }
  //#endregion

  //#region Color
  if (strcmp(cmd, "c") == 0) {
    // meh, doesnt work otherwise for some reason
    char str[5 ];
    snprintf(str, 5, "0x%02X", ws2812fx.getColor() >> 16);
    Serial.print(str);
    char str2[5];
    snprintf(str2, 5, "%04X", ws2812fx.getColor());
    Serial.println(str2);

    isoke = true;
    hasFullAnswer = true;
  }

  if (strncmp(cmd, "c ", 2) == 0) {
    currentState.Color = (uint32_t)strtoul(cmd + 2, NULL, 16);
    isoke = true;
  }
  
  if (strncmp(cmd, "sp ", 3) == 0) {
    if (!previousState.Sparking) {
      previousState = currentState;
    }
    
    currentState = State {
      (uint32_t)strtoul(cmd + 3, NULL, 16),
      255,
      FX_MODE_STATIC,
      100,
      true,
      millis(),
      currentState.SparkDuration,

      false,
      previousState.Brightness,
    };
    
    isoke = true;
  }
  //#endregion

  if (strcmp(cmd, "r") == 0) {
    reboot();
  }

  if (strcmp(cmd, "state") == 0) {
    printState(currentState);
    hasFullAnswer = true;
  }

  if (strcmp(cmd, "dbg") == 0) {
    printState(previousState);
    printState(currentState);
    hasFullAnswer = true;
  }

  if (isoke) {
      setState(currentState);
  }

  if (!hasFullAnswer) {
    if (isoke) {
      Serial.println("OK");
    } else {
      Serial.println("NOT OK");
    }
  }

  cmd[0] = '\0';
  cmd_complete = false;
}


/*
 * Reads new input from serial to cmd string. Command is completed on \n
 */
void recvChar(void) {
  static byte index = 0;
  while (Serial.available() > 0 && cmd_complete == false) {
    char rc = Serial.read();
    if (rc != '\n') {
      if(index < MAX_NUM_CHARS) cmd[index++] = rc;
    } else {
      cmd[index] = '\0'; // terminate the string
      index = 0;
      cmd_complete = true;
    }
  }
}
