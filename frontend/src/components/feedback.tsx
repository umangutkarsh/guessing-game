// components/feedback.tsx
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
		<div className='p-4 bg-gray-800 rounded-md'>
			{correct ? (
				<h2 className='text-2xl font-bold text-green-400'>Correct! 🎉</h2>
			) : (
				<h2 className='text-2xl font-bold text-red-400'>Incorrect! 😢</h2>
			)}
			<p className='mt-2'>
				<span className='font-semibold'>Fun Fact:</span> {funFact}
			</p>
			<p className='mt-1'>
				<span className='font-semibold'>Trivia:</span> {trivia}
			</p>
			<p className='mt-2 text-lg'>Your Score: {score}</p>
		</div>
	);
};

export default Feedback;
