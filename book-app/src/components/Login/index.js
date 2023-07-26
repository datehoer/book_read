import React, { useState } from "react"
import ApiService from "../../api"
import { useNavigate } from "react-router-dom"
import Cookies from 'js-cookie'
import md5 from 'js-md5'
import "./index.css"

function Login() {
    const navigate = useNavigate()
    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")

    const handleLogin = async (e) => {
        e.preventDefault()
        try {
            await ApiService.login(username, password)
            // 创建一个随机字符串
            const randomString = Math.random().toString(36).substring(2, 15) + Math.random().toString(36).substring(2, 15);

            // 创建一个MD5 hash
            const hash = md5(randomString);

            // 设置cookie
            Cookies.set('auth', hash, { expires: 1 }); // expires after 1 day
            navigate("/")
        } catch (error) {
            if (error.response) {
                // The request was made and the server responded with a status code
                // that falls out of the range of 2xx
                alert(error.response.data.message)
            } else {
                // Something happened in setting up the request that triggered an Error
                alert("Error: " + error.message)
            }
        }
    }

    return (
        <div className='login-container'>
            <form className='login-form' onSubmit={handleLogin}>
                <h2>书库登录</h2>
                <div className='input-group'>
                    <label htmlFor='username'>用户名：</label>
                    <input
                        type='text'
                        id='username'
                        value={username}
                        placeholder='Username'
                        onChange={(e) => setUsername(e.target.value)}
                        required
                    />
                </div>
                <div className='input-group'>
                    <label htmlFor='password'>密码：</label>
                    <input
                        type='password'
                        id='password'
                        placeholder='Password'
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        required
                    />
                </div>
                <button type='submit' className='login-button'>
                    Login
                </button>
            </form>
        </div>
    )
}

export default Login
