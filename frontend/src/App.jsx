import './index.css'
import { HashRouter as Router, Route, Routes } from 'react-router-dom'
import MainPage from './Pages/Home'
import ProductPage from './Pages/Products';

function App() {
    return (
        <Router>
            <Routes>
                <Route path="/" element={<MainPage/>}/>
                <Route path="/products" element={<ProductPage/>}/>
            </Routes>
        </Router>
    );
}

export default App
