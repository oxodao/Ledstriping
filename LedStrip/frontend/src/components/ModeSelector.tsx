import { useLedStrip } from "../hooks/StripState";
import Card from "./Card";
import ListItem from "./ListItem";

import '../assets/css/list.scss';

// @TODO: get this from the DB through /api/state
const modes = [
    'Static',
    'Blink',
    'Breath',
    'Color Wipe',
    'Color Wipe Inverse',
    'Color Wipe Reverse',
    'Color Wipe Reverse Inverse',
    'Color Wipe Random',
    'Random Color',
    'Single Dynamic',
    'Multi Dynamic',
    'Rainbow',
    'Rainbow Cycle',
    'Scan',
    'Dual Scan',
    'Fade',
    'Theater Chase',
    'Theater Chase Rainbow',
    'Running Lights',
    'Twinkle',
    'Twinkle Random',
    'Twinkle Fade',
    'Twinkle Fade Random',
    'Sparkle',
    'Flash Sparkle',
    'Hyper Sparkle',
    'Strobe',
    'Strobe Rainbow',
    'Multi Strobe',
    'Blink Rainbow',
    'Chase White',
    'Chase Color',
    'Chase Random',
    'Chase Rainbow',
    'Chase Flash',
    'Chase Flash Random',
    'Chase Rainbow White',
    'Chase Blackout',
    'Chase Blackout Rainbow',
    'Color Sweep Random',
    'Running Color',
    'Running Red Blue',
    'Running Random',
    'Larson Scanner',
    'Comet',
    'Fireworks',
    'Fireworks Random',
    'Merry Christmas',
    'Fire Flicker',
    'Fire Flicker (soft)',
    'Fire Flicker (intense)',
    'Circus Combustus',
    'Halloween',
    'Bicolor Chase',
    'Tricolor Chase',
    'TwinkleFOX',
    'Custom 0',
    'Custom 1',
    'Custom 2',
    'Custom 3',
    'Custom 4',
    'Custom 5',
    'Custom 6',
    'Custom 7',
]

export default function ModeSelector() {
    const strip = useLedStrip();

    return <Card title="Mode">
        <ul className="list">
            {
                modes.map((val, i) => <ListItem text={val} value={i} selected={strip.Mode === i} action={() => strip.setValue('Mode', i)}/>)
            }
        </ul>
    </Card>
}