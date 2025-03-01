'use client';
import { useState } from 'react';
import api from '@/lib/api';
import axios from 'axios';
import Image from 'next/image';

export default function Challenge() {
	const [inviteData, setInviteData] = useState<{
		challengeId: string;
		inviteLink: string;
		inviteImage: string;
	} | null>(null);
	const [error, setError] = useState<string | null>(null);
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

	return (
		<div className='min-h-screen bg-gray-900 text-white px-4 py-6'>
			<h1 className='text-3xl font-bold mb-4'>Challenge a Friend</h1>
			<button
				onClick={handleCreateChallenge}
				className='bg-blue-600 hover:bg-blue-700 px-4 py-2 rounded font-semibold'
			>
				Create Challenge Invite
			</button>
			{error && <p className='text-red-500 mt-4'>{error}</p>}
			{inviteData && (
				<div className='mt-6 space-y-4'>
					<p className='text-lg'>Share this invite link on WhatsApp:</p>
					<p className='break-words'>
						<a
							href={`${encodeURIComponent(inviteData.inviteLink)}`}
							className='text-blue-400 underline'
						>
							{inviteData.inviteLink}
						</a>
					</p>
					<div className='w-[300px]'>
						<Image
							src={inviteData.inviteImage}
							alt='Challenge Invite'
							width={300}
							height={300}
							className='rounded'
						/>
					</div>
				</div>
			)}
		</div>
	);
}
