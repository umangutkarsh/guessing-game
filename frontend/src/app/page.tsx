'use client';
import { useState, useEffect } from 'react';
import api from '@/lib/api';
import { useRouter } from 'next/navigation';
import axios from 'axios';

export default function Home() {
	const [username, setUsername] = useState('');
	const [error, setError] = useState<string | null>(null);
	const [registeredUser, setRegisteredUser] = useState<{
		userId: string;
		username: string;
	} | null>(null);
	const router = useRouter();

	// Check if a user is already registered on mount.
	useEffect(() => {
		const storedUserId = localStorage.getItem('userId');
		const storedUsername = localStorage.getItem('username');
		if (storedUserId && storedUsername) {
			setRegisteredUser({ userId: storedUserId, username: storedUsername });
		}
	}, []);

	const handleRegister = async (e: React.FormEvent) => {
		e.preventDefault();
		setError(null);
		try {
			const res = await api.post('/auth/register', { username });
			// Store both userId and username in localStorage.
			localStorage.setItem('userId', res.data.userId);
			localStorage.setItem('username', res.data.username);
			setRegisteredUser({
				userId: res.data.userId,
				username: res.data.username,
			});
			// Optionally, you can navigate to the game immediately:
			// router.push('/game');
		} catch (err: unknown) {
			if (axios.isAxiosError(err)) {
				setError(err.response?.data?.error || 'Registration failed');
			} else {
				setError('Registration failed');
			}
		}
	};

	// Continue with the registered user and go to the game.
	const handlePlay = () => {
		router.push('/game');
	};

	// Clear the stored user data and allow new registration.
	const handleRegisterNewUser = () => {
		localStorage.removeItem('userId');
		localStorage.removeItem('username');
		setRegisteredUser(null);
		setUsername('');
	};

	return (
		<div className='min-h-screen flex flex-col items-center justify-center bg-gray-900 text-white px-4'>
			<h1 className='text-3xl font-bold mb-6'>Welcome to Globetrotter</h1>
			{registeredUser ? (
				<div className='space-y-4 w-full max-w-sm'>
					<p className='text-lg'>
						You are logged in as{' '}
						<span className='font-semibold'>{registeredUser.username}</span>.
					</p>
					<button
						onClick={handlePlay}
						className='bg-blue-600 hover:bg-blue-700 px-4 py-2 rounded font-semibold'
					>
						Play as {registeredUser.username}
					</button>
					<button
						onClick={handleRegisterNewUser}
						className='bg-red-600 hover:bg-red-700 px-4 py-2 rounded font-semibold'
					>
						Register New User
					</button>
				</div>
			) : (
				<form onSubmit={handleRegister} className='space-y-4 w-full max-w-sm'>
					<label className='block'>
						<span className='block mb-2'>Enter a unique username:</span>
						<input
							type='text'
							value={username}
							onChange={e => setUsername(e.target.value)}
							required
							className='w-full p-2 rounded bg-gray-800 focus:outline-none focus:ring-2 focus:ring-blue-500'
						/>
					</label>
					<button
						type='submit'
						className='bg-blue-600 hover:bg-blue-700 px-4 py-2 rounded font-semibold'
					>
						Register
					</button>
				</form>
			)}
			{error && <p className='text-red-500 mt-4'>{error}</p>}
		</div>
	);
}
