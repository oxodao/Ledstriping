import '../assets/css/list.scss';

type ListItemProps = {
    text: string
    value: number
    selected: boolean,
    action: (value: number) => void
}

export default function ListItem({text, value, selected, action}: ListItemProps) {
    return <li className={`Item${selected ? ' selected' : ''}`} onClick={() => action(value)}>{text}</li>
}