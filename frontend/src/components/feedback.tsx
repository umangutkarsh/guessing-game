import React from 'react';

interface FeedbackProps {
	correct: boolean;
	funFact: string;
	trivia: string;
	score: number;
}

const Feedback: React.FC<FeedbackProps> = ({
	correct,
	funFact,
	trivia,
	score,
}) => {
	return (
		<div className='card'>
			{correct ? (
				<h2 style={{ color: 'green' }}>Correct! 🎉</h2>
			) : (
				<h2 style={{ color: 'red' }}>Incorrect! 😢</h2>
			)}
			<p>
				<strong>Fun Fact:</strong> {funFact}
			</p>
			<p>
				<strong>Trivia:</strong> {trivia}
			</p>
			<p>Your Score: {score}</p>
		</div>
	);
};

export default Feedback;
