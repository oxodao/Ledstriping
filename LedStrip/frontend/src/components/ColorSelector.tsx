import Card from "./Card";
import { ChromePicker } from 'react-color';
import { useLedStrip } from "../hooks/StripState";

export default function ColorSelector() {
    const strip = useLedStrip();

    return <Card title="Color">
        <ChromePicker color={strip.Color} onChange={(col) => strip.setValue('Color', col.hex)} disableAlpha/>
    </Card>
}