import axios from 'axios';
import Api from '../api';

jest.mock('axios');
const mockedAxios = axios as jest.Mocked<typeof axios>;

const BASE_API_URL: string = 'http://localhost:8080';

describe('should work with axios mock', () => {
	it('calls axios and returns a task array', async () => {
		mockedAxios.get.mockImplementationOnce(() =>
			Promise.resolve({
				data: [
					{
						_id: 'test_01',
						description: 'task_01',
						status: true
					},
					{
						_id: 'test_02',
						description: 'task_02',
						status: false
					}
				]
			})
		);

		const items = await Api.GetAllTask();

		expect(items.data).toEqual([
			{
				_id: 'test_01',
				description: 'task_01',
				status: true
			},
			{
				_id: 'test_02',
				description: 'task_02',
				status: false
			}
		]);
		expect(mockedAxios.get).toHaveBeenCalledTimes(1);
		expect(mockedAxios.get).toHaveBeenCalledWith(BASE_API_URL + '/api/task');
	});

	it('calls axios and creates a task', async () => {
		mockedAxios.post.mockImplementationOnce(() =>
			Promise.resolve({
				data: {
					_id: 'test_03',
					description: 'task_03',
				}
			})
		);

		const items = await Api.CreateTask('task_03');

		expect(items.data).toEqual({
			_id: 'test_03',
			description: 'task_03',
		});
		expect(mockedAxios.post).toHaveBeenCalledTimes(1);
		expect(mockedAxios.post).toHaveBeenCalledWith(
			BASE_API_URL + '/api/task',
			{ description: 'task_03' },
			{
				headers: {
					'Content-Type': 'application/x-www-form-urlencoded'
				}
			}
		);
	});

	it('calls axios and completes a task', async () => {
		mockedAxios.put.mockImplementationOnce(() =>
			Promise.resolve({
				data: { _id: 'test_04' }
			})
		);

		const items = await Api.CompleteTask('test_04');

		expect(items.data).toEqual({ _id: 'test_04' });
		expect(mockedAxios.put).toHaveBeenCalledTimes(1);
		expect(mockedAxios.put).toHaveBeenCalledWith(
			BASE_API_URL + '/api/task/' + 'test_04',
			{
				headers: {
					'Content-Type': 'application/x-www-form-urlencoded'
				}
			}
		);
	});
});
