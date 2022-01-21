import Card from "./Card";
import { StripNumberProp, useLedStrip } from "../hooks/StripState";
import Slider from 'rc-slider';

import 'rc-slider/assets/index.css';

type SliderCardProps = {
    title: string,
    prop: StripNumberProp,
    valueCalculator: (val: number) => string,
    min: number,
    max: number,
    step?: number,
}

export default function SliderCard({title, prop, valueCalculator, min, max, step}: SliderCardProps) {
    const strip = useLedStrip();

    console.log(prop, strip)

    return <Card title={title} className="Slider">
        <h2>{valueCalculator(strip[prop])}</h2>
        <Slider
            className="slider"
            min={min}
            max={max}
            marks={{0: '0%', [max * .25]: '25%', [max * .5]: '50%', [max * .75]: '75%', [max]: '100%'}}
            step={step}
            included={false}
            value={strip[prop]}
            onChange={(val) => strip.setValue(prop, val)}
        />
    </Card>
}

export function BrightnessSelector() {
    return <SliderCard title="Brightness" min={0} max={255} prop="Brightness" valueCalculator={(val) => {
        console.log(val);
        console.log(val/255);
        console.log(val/255*100);
        console.log(Math.floor(val/255*100));
        return Math.floor(val / 255 * 100) + "%";
    }}/>
}

export function SpeedSelector() {
    return <SliderCard title="Speed" prop="Speed" min={0} max={10000} valueCalculator={val => val + "ms"} step={500}/>
}