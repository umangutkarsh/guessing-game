import React from 'react';

interface ScoreBoardProps {
	correct: number;
	incorrect: number;
}

const ScoreBoard: React.FC<ScoreBoardProps> = ({ correct, incorrect }) => {
	return (
		<div className='card' style={{ display: 'inline-block' }}>
			<h3>Correct: {correct}</h3>
			<h3>Incorrect: {incorrect}</h3>
		</div>
	);
};

export default ScoreBoard;
