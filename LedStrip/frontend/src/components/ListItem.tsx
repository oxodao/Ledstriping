import '../assets/css/list.scss';

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faEdit, faTrash } from '@fortawesome/free-solid-svg-icons'

type ListItemProps = {
    text: string
    selected?: boolean,
    action: () => void

    editAction?: () => void,
    deleteAction?: () => void,
}

export default function ListItem({text, selected, action, editAction, deleteAction}: ListItemProps) {
    return <li className={`Item${selected ? ' selected' : ''}`}>
        <span onClick={() => action()}>{text}</span>
        {
            editAction && <FontAwesomeIcon icon={faEdit} onClick={editAction}/>
        }
        {
            deleteAction && <FontAwesomeIcon icon={faTrash} onClick={deleteAction}/>
        }
    </li>
}