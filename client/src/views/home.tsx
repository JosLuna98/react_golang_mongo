import React, { useEffect, useState } from 'react';
import { Card, Form, Input, Button } from 'semantic-ui-react';
import logo from '../logo.svg';
import Task from '../models/task';
import Api from '../services/api';
import TaskCard from '../components/task_card';

const Home = () => {
	const [ task, setTask ]: [string, (task: string) => void] = useState('');
	const [ items, setItems ]: [Array<Task>, (items: Array<Task>) => void] = useState([] as Array<Task>);

	const getTasks: () => void = () => {
		Api.GetAllTask().then((res) => {
			const result: Array<Task> = res.data;
			setItems(result);
		});
	};

	useEffect(() => getTasks(), []);

	const renderItems = () => {
		if (!items) {
			return (
				<Card fluid>
					<Card.Content>
						<Card.Header textAlign="center">
							<div>There're no tasks</div>
						</Card.Header>
					</Card.Content>
				</Card>
			);
		} else {
			return items.map((taskItem) => <TaskCard taskItem={taskItem} getTasks={getTasks}/>);
		}
	};

	const onSubmit = () => {
		if (task) {
			Api.CreateTask(task).then((_) => getTasks()).finally(() => setTask(''));
		}
	};

	const onChange = (event: React.ChangeEvent<HTMLInputElement>) => setTask(event.target.value);

	return (
		<div>
			<div className="row">
				<img src={logo} className="App-logo" alt="logo" />
			</div>
			<div className="row">
				<Form onSubmit={onSubmit}>
					<Input type="text" name="task" onChange={onChange} value={task} fluid placeholder="Task description" />
					<Button>Create Task</Button>
				</Form>
			</div>
			<div className="row">
				<Card.Group>{renderItems()}</Card.Group>
			</div>
		</div>
	);
};

export default Home;
