// components/destination-options.tsx
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
		<div className='p-4 bg-gray-800 rounded-md'>
			<h2 className='text-xl font-semibold mb-2'>Choose a Destination</h2>
			<ul className='space-y-2'>
				{options.map(option => (
					<li key={option.cityId}>
						<label className='flex items-center space-x-2'>
							<input
								type='radio'
								name='destination'
								value={option.cityId}
								checked={selectedOption === option.cityId}
								onChange={() => onOptionSelect(option.cityId)}
								className='form-radio h-4 w-4 text-blue-600'
							/>
							<span>{option.cityName}</span>
						</label>
					</li>
				))}
			</ul>
		</div>
	);
};

export default DestinationOptions;
