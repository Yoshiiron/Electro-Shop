import './Header.css'

function Header() {
    return (
        <header>
            <div className='header-container'>
                <nav className='nav-links'>
                    <a>ELECTRO-SHOP</a>
                    <a href="/">Home</a>
                    <a href="/#/products">Products</a>
                    <a href="">About</a>
                </nav>
            </div>
        </header>
    );
}

export default Header;