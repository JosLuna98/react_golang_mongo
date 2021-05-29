import axios from 'axios';

const BASE_API_URL: string = 'http://localhost:8080';

const Api = {
	GetAllTask: () => axios.get(BASE_API_URL + '/api/task'),

	CreateTask: (description: string) =>
		axios.post(
			BASE_API_URL + '/api/task',
			{ description },
			{
				headers: {
					'Content-Type': 'application/x-www-form-urlencoded'
				}
			}
		),

	CompleteTask: (id: string) =>
		axios.put(BASE_API_URL + '/api/task/' + id, {
			headers: {
				'Content-Type': 'application/x-www-form-urlencoded'
			}
		}),

	UndoTask: (id: string) =>
		axios.put(BASE_API_URL + '/api/undoTask/' + id, {
			headers: {
				'Content-Type': 'application/x-www-form-urlencoded'
			}
		}),

	DeleteTask: (id: string) =>
		axios.delete(BASE_API_URL + '/api/deleteTask/' + id, {
			headers: {
				'Content-Type': 'application/x-www-form-urlencoded'
			}
		})
};

export default Api;
