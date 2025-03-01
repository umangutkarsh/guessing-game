// components/clue-display.tsx
import React from 'react';

interface ClueDisplayProps {
	clues: string[];
}

const ClueDisplay: React.FC<ClueDisplayProps> = ({ clues }) => {
	return (
		<div className='p-4 bg-gray-800 rounded-md'>
			<h2 className='text-xl font-semibold mb-2'>Clues</h2>
			<ul className='list-disc ml-5 space-y-1'>
				{clues.map((clue, index) => (
					<li key={index}>{clue}</li>
				))}
			</ul>
		</div>
	);
};

export default ClueDisplay;
