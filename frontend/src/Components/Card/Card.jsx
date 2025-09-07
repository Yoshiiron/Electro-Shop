import './Card.css'

function Card() {
    return (
        <div className="card">
            <img className="card-image" src="https://dlcdnwebimgs.asus.com/files/media/B54349DD-1E57-46B1-B1A7-10166BF72325/v1/img/gaming/pd.jpg" alt="Product" />
            <h2>Product Title</h2>
            <p>$99.99</p>
            <button>Add to Cart</button>
        </div>
    );
}

export default Card;