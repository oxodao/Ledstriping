import { useLedStrip } from "../hooks/StripState";
import Card from "./Card";
import ListItem from "./ListItem";
import { useData } from "../hooks/DataState";

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faPlus } from '@fortawesome/free-solid-svg-icons'

import '../assets/css/list.scss';
import '../assets/css/favorite_browser.scss';

export default function FavoriteBrowser() {
    const data = useData();
    const strip = useLedStrip();

    const addFavorite = () => {
        data.addFavorite({
            "Name": data.currentName,
            "Color": strip.Color,
            "Mode": strip.Mode,
            "Brightness": strip.Brightness,
            "Speed": strip.Speed,
        })
    }

    return <Card title="Favorite">
        <ul className="list">
            {
                data.Favorites.map((val, i) => (
                    <ListItem
                        key={i}
                        text={val.Name}
                        action={() => strip.useFavorite(val)}
                        editAction={() => { console.log("Edit " + val.Name) }}
                        deleteAction={() => { console.log("Delete " + val.Name) }}
                    />
                ))
            }
        </ul>
        <div className="Add">
            <input className="Add__Field" placeholder="Name" type="text" value={data.currentName} onChange={(name) => data.setCurrentName(name.target.value)} />
            <FontAwesomeIcon className="Add__Button" icon={faPlus} onClick={addFavorite} />
        </div>
    </Card>
}