import { useLedStrip } from "../hooks/StripState";
import Card from "./Card";
import ListItem from "./ListItem";

import '../assets/css/list.scss';
import { useData } from "../hooks/DataState";

export default function ModeSelector() {
    const data = useData();
    const strip = useLedStrip();

    return <Card title="Mode">
        <ul className="list">
            {
                data.Modes.map((val, i) => <ListItem key={i} text={val} selected={strip.Mode === i} action={() => strip.setValue('Mode', i)}/>)
            }
        </ul>
    </Card>
}