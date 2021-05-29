import './app.css';
import { Container } from 'semantic-ui-react';
import Home from './views/home';

const App = () => {
	return (
		<div>
			<Container>
				<Home />
			</Container>
		</div>
	);
}

export default App;
