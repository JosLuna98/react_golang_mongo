import { render, screen } from '@testing-library/react';
import App from './app';
import TaskCard from './components/task_card';

describe("shoulds render some Card Component", () => {
  it(`must display "there're no tasks"`, () => {
    render(<App />);
    expect(screen.queryByText(/there're no tasks/i)).toBeInTheDocument();
  });

  it("must display task description", () => {
    const taskItem = {
      _id: "test_01",
      description: "task_01",
      status: true
    };

    render(<TaskCard taskItem={taskItem} getTasks={()=>{}}/>);
    expect(screen.queryByText(/task_01/i)).toBeInTheDocument();
  });
});