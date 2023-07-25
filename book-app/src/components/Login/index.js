import React, { useState } from 'react';
import axiosInstance from '../../api';

function Login() {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    const handleLogin = async () => {
        try {
            const response = await axiosInstance.post('/api/login', {
                username,
                password
            });
            console.log(response)
        } catch (error) {
            if (error.response && error.response.status === 401) {
                // handle 401 error here, for example:
                alert('Unauthorized. Please check your username or password.');
            }
        }
    };

    return (
        <div>
            <input 
                type="text" 
                value={username} 
                onChange={e => setUsername(e.target.value)} 
                placeholder="Username"
            />
            <input 
                type="password" 
                value={password} 
                onChange={e => setPassword(e.target.value)} 
                placeholder="Password"
            />
            <button onClick={handleLogin}>Login</button>
        </div>
    );
}

export default Login;