import { Card, Icon, SemanticCOLORS } from 'semantic-ui-react';
import Task from '../models/task';
import Api from '../services/api';
import { VoidFunction, onResponse } from '../utils/common';

const TaskCard = ({ taskItem, getTasks }: { taskItem: Task; getTasks: VoidFunction }) => {
	let color: SemanticCOLORS = taskItem.status ? 'green' : 'yellow';
	let style: React.CSSProperties = {
		wordWrap: 'break-word',
		textDecorationLine: taskItem.status ? 'line-through' : ''
	};

	const completeTask = (id: string) => Api.CompleteTask(id).then((res) => onResponse(res, getTasks));

	const undoTask = (id: string) => Api.UndoTask(id).then((res) => onResponse(res, getTasks));

	const deleteTask = (id: string) => Api.DeleteTask(id).then((res) => onResponse(res, getTasks));

	return (
		<Card key={taskItem._id} color={color} fluid>
			<Card.Content>
				<Card.Header textAlign="left">
					<div style={style}>{taskItem.description}</div>
				</Card.Header>
				<Card.Meta textAlign="right">
					<Icon name="check circle" color="green" onClick={() => completeTask(taskItem._id)} />
					<span style={{ paddingRight: 10 }}>Done</span>
					<Icon name="undo" color="yellow" onClick={() => undoTask(taskItem._id)} />
					<span style={{ paddingRight: 10 }}>Undo</span>
					<Icon name="delete" color="red" onClick={() => deleteTask(taskItem._id)} />
					<span style={{ paddingRight: 10 }}>Delete</span>
				</Card.Meta>
			</Card.Content>
		</Card>
	);
};

export default TaskCard;
