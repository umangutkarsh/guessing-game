'use client';
import { useState, useEffect } from 'react';
import api from '@/lib/api';
import { useRouter } from 'next/navigation';
import axios from 'axios';

export default function Home() {
	const [username, setUsername] = useState('');
	const [error, setError] = useState<string | null>(null);
	// Registered user data if exists
	const [registeredUser, setRegisteredUser] = useState<{
		userId: string;
		username: string;
	} | null>(null);
	// Toggle between Registration and Login forms
	const [isLogin, setIsLogin] = useState<boolean>(false);
	const router = useRouter();

	// On mount, check if user data exists in localStorage.
	useEffect(() => {
		const storedUserId = localStorage.getItem('userId');
		const storedUsername = localStorage.getItem('username');
		if (storedUserId && storedUsername) {
			setRegisteredUser({ userId: storedUserId, username: storedUsername });
		}
	}, []);

	// Registration handler
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
			// Optionally, navigate to the game immediately:
			// router.push('/game');
		} catch (err: unknown) {
			if (axios.isAxiosError(err)) {
				setError(err.response?.data?.error || 'Registration failed');
			} else {
				setError('Registration failed');
			}
		}
	};

	// Login handler – assumes a backend endpoint /auth/login that returns the user profile.
	const handleLogin = async (e: React.FormEvent) => {
		e.preventDefault();
		setError(null);
		try {
			const res = await api.get(`/auth/profile/username/${username}`);
			// Store both userId and username in localStorage.
			localStorage.setItem('userId', res.data.userId);
			localStorage.setItem('username', res.data.username);
			setRegisteredUser({
				userId: res.data.userId,
				username: res.data.username,
			});
		} catch (err: unknown) {
			if (axios.isAxiosError(err)) {
				setError(err.response?.data?.error || 'Login failed');
			} else {
				setError('Login failed');
			}
		}
	};

	// Proceed to game if user is registered
	const handlePlay = () => {
		router.push('/game');
	};

	// Clear user data to allow new registration
	const handleRegisterNewUser = () => {
		localStorage.removeItem('userId');
		localStorage.removeItem('username');
		setRegisteredUser(null);
		setUsername('');
		setIsLogin(false);
	};

	return (
		<div className='container'>
			<h1 className='heading'>Welcome to Globetrotter</h1>
			{registeredUser ? (
				<div>
					<p>
						You are logged in as <strong>{registeredUser.username}</strong>.
					</p>
					<button onClick={handlePlay} className='button'>
						Play as {registeredUser.username}
					</button>
					<button
						onClick={handleRegisterNewUser}
						className='button button-secondary'
					>
						Register New User
					</button>
				</div>
			) : (
				<div>
					<div>
						<button onClick={() => setIsLogin(false)} className='button'>
							Register
						</button>
						<button
							onClick={() => setIsLogin(true)}
							className='button button-secondary'
						>
							Login
						</button>
					</div>
					{isLogin ? (
						<form onSubmit={handleLogin}>
							<div className='form-group'>
								<label className='label'>Enter your username:</label>
								<input
									type='text'
									value={username}
									onChange={e => setUsername(e.target.value)}
									required
									className='input'
								/>
							</div>
							<button type='submit' className='button'>
								Login
							</button>
						</form>
					) : (
						<form onSubmit={handleRegister}>
							<div className='form-group'>
								<label className='label'>Enter a unique username:</label>
								<input
									type='text'
									value={username}
									onChange={e => setUsername(e.target.value)}
									required
									className='input'
								/>
							</div>
							<button type='submit' className='button'>
								Register
							</button>
						</form>
					)}
				</div>
			)}
			{error && <p className='error'>{error}</p>}
		</div>
	);
}
