'use client';
import { useState } from 'react';
import api from '@/lib/api';
import axios from 'axios';
import Image from 'next/image';
import { useRouter } from 'next/navigation';

export default function Challenge() {
	const [inviteData, setInviteData] = useState<{
		challengeId: string;
		inviteLink: string;
		inviteImage: string;
	} | null>(null);
	const [error, setError] = useState<string | null>(null);
	const router = useRouter();
	const userId =
		typeof window !== 'undefined' ? localStorage.getItem('userId') : null;

	const handleCreateChallenge = async () => {
		if (!userId) {
			setError('User not logged in.');
			return;
		}
		try {
			const res = await api.post('/challenge', { userId });
			setInviteData(res.data);
		} catch (err: unknown) {
			if (axios.isAxiosError(err)) {
				setError(err.response?.data?.error || 'Error creating challenge');
			} else {
				setError('Error creating challenge');
			}
		}
	};

	// New function to navigate back to the main page
	const goToMainPage = () => {
		router.push('/');
	};

	return (
		<div className='container'>
			<h1 className='heading'>Challenge a Friend</h1>
			<div style={{ marginBottom: '1rem' }}>
				<button onClick={goToMainPage} className='button button-secondary'>
					Back to Main
				</button>
			</div>
			<button onClick={handleCreateChallenge} className='button'>
				Create Challenge Invite
			</button>
			{error && <p style={{ color: 'red' }}>{error}</p>}
			{inviteData && (
				<div className='card'>
					<p>Share this invite link on WhatsApp:</p>
					<p>
						<a href={inviteData.inviteLink} className='link'>
							{inviteData.inviteLink}
						</a>
					</p>
					<div
						style={{ width: '300px', position: 'relative', height: '300px' }}
					>
						<Image
							src={inviteData.inviteImage}
							alt='Challenge Invite'
							fill
							className='rounded'
							style={{ objectFit: 'cover' }}
						/>
					</div>
				</div>
			)}
		</div>
	);
}
