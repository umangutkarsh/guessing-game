import React from 'react';

interface ClueDisplayProps {
	clues: string[];
}

const ClueDisplay: React.FC<ClueDisplayProps> = ({ clues }) => {
	return (
		<div className='card'>
			<h2>Clues</h2>
			<ul className='list'>
				{clues.map((clue, index) => (
					<li key={index}>{clue}</li>
				))}
			</ul>
		</div>
	);
};

export default ClueDisplay;
