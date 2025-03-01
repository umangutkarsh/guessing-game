import React from 'react';

interface ScoreBoardProps {
	correct: number;
	incorrect: number;
}

const ScoreBoard: React.FC<ScoreBoardProps> = ({ correct, incorrect }) => {
	return (
		<div className='p-2 bg-gray-800 inline-block rounded-md'>
			<h3 className='font-semibold text-lg'>Correct: {correct}</h3>
			<h3 className='font-semibold text-lg'>Incorrect: {incorrect}</h3>
		</div>
	);
};

export default ScoreBoard;
