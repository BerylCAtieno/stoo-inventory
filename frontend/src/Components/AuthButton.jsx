import '../CSS/authButton.css'

function AuthButton({label, onClick}) {
    return (
        <button className="auth-button" onClick={onClick}>
            {label}
        </button>
    )
}

export default AuthButton