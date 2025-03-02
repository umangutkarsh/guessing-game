import React from 'react';

interface Option {
	cityId: string;
	cityName: string;
}

interface DestinationOptionsProps {
	options: Option[];
	selectedOption: string;
	onOptionSelect: (cityId: string) => void;
}

const DestinationOptions: React.FC<DestinationOptionsProps> = ({
	options,
	selectedOption,
	onOptionSelect,
}) => {
	return (
		<div className='card'>
			<h2>Choose a Destination</h2>
			<ul>
				{options.map(option => (
					<li key={option.cityId}>
						<label>
							<input
								type='radio'
								name='destination'
								value={option.cityId}
								checked={selectedOption === option.cityId}
								onChange={() => onOptionSelect(option.cityId)}
							/>
							{option.cityName}
						</label>
					</li>
				))}
			</ul>
		</div>
	);
};

export default DestinationOptions;
