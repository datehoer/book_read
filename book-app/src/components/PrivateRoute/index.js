import { useEffect } from "react"
import { useNavigate } from "react-router-dom"
import Cookies from "js-cookie"

function PrivateRoute({ children }) {
    let navigate = useNavigate()
    let isLoggedIn = !!Cookies.get("auth")

    useEffect(() => {
        if (!isLoggedIn) {
            navigate("/login")
        }
    }, [isLoggedIn, navigate])

    return isLoggedIn ? children : null
}

export default PrivateRoute
