'use client';
import { useEffect, useState } from 'react';
import api from '@/lib/api';
import ClueDisplay from '@/components/clue-display';
import DestinationOptions from '@/components/destination-options';
import Feedback from '@/components/feedback';
import ScoreBoard from '@/components/score-board';
import { useRouter } from 'next/navigation';

type DestinationResponse = {
	questionToken: string;
	clues: string[];
	options: { cityId: string; cityName: string }[];
};

type GuessResult = {
	correct: boolean;
	funFact: string;
	trivia: string;
	score: number; // backend returned correct count (optional)
};

export default function Game() {
	const [destination, setDestination] = useState<DestinationResponse | null>(
		null
	);
	const [selectedOption, setSelectedOption] = useState<string>('');
	const [result, setResult] = useState<GuessResult | null>(null);
	const [loading, setLoading] = useState<boolean>(false);
	const [correctCount, setCorrectCount] = useState<number>(0);
	const [incorrectCount, setIncorrectCount] = useState<number>(0);

	const router = useRouter();
	const userId =
		typeof window !== 'undefined' ? localStorage.getItem('userId') : null;

	const fetchDestination = async () => {
		setLoading(true);
		try {
			const res = await api.get<DestinationResponse>('/game/destination');
			setDestination(res.data);
			setResult(null);
			setSelectedOption('');
		} catch (err) {
			console.error(err);
		} finally {
			setLoading(false);
		}
	};

	const submitGuess = async () => {
		if (!destination || !userId) return;
		try {
			const res = await api.post<GuessResult>('/game/guess', {
				userId: userId,
				questionToken: destination.questionToken,
				selectedCityId: selectedOption,
			});
			setResult(res.data);
			if (res.data.correct) {
				setCorrectCount(prev => prev + 1);
			} else {
				setIncorrectCount(prev => prev + 1);
			}
		} catch (err) {
			console.error(err);
		}
	};

	const handleChangeUser = () => {
		// Clear the stored userId (and optionally username) and redirect to the registration page.

		router.push('/');
	};

	useEffect(() => {
		if (!userId) {
			router.push('/');
		} else {
			fetchDestination();
		}
	}, [userId, router]);

	return (
		<div className='min-h-screen bg-gray-900 text-white px-4 py-6'>
			<div className='flex justify-between items-center mb-4'>
				<h1 className='text-3xl font-bold'>Globetrotter Challenge</h1>
				<button
					onClick={handleChangeUser}
					className='bg-red-600 hover:bg-red-700 px-4 py-2 rounded font-semibold'
				>
					Register New User
				</button>
			</div>
			{/* ScoreBoard now displays both correct and incorrect counts */}
			<ScoreBoard correct={correctCount} incorrect={incorrectCount} />
			{loading && <p className='mt-4'>Loading...</p>}
			{destination && !result && (
				<div className='mt-4 space-y-4'>
					<ClueDisplay clues={destination.clues} />
					<DestinationOptions
						options={destination.options}
						selectedOption={selectedOption}
						onOptionSelect={setSelectedOption}
					/>
					<button
						onClick={submitGuess}
						disabled={!selectedOption}
						className='bg-blue-600 hover:bg-blue-700 px-4 py-2 rounded font-semibold'
					>
						Submit Guess
					</button>
				</div>
			)}
			{result && (
				<div className='mt-6'>
					<Feedback {...result} />
					<button
						onClick={fetchDestination}
						className='mt-4 bg-green-600 hover:bg-green-700 px-4 py-2 rounded font-semibold'
					>
						Play Again
					</button>
				</div>
			)}
		</div>
	);
}
