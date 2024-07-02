import express from 'express';
import dotenv from 'dotenv';
import Stripe from 'stripe';
import mysql from 'mysql';
import axios from 'axios';

// Load environment variables
dotenv.config();

const app = express();
app.use(express.static('public'));
app.use(express.json());

// Middleware for logging request body
app.use((req, res, next) => {
    console.log("Request body:", req.body);
    next();
});

// Database connection
const db = mysql.createConnection({
    host: 'localhost',
    port: 3306,
    user: 'root',
    password: '',
    database: 'onnlinefood'
});

db.connect(err => {
    if (err) {
        console.error('Error connecting to database:', err);
    } else {
        console.log('Connected to database');
    }
});

app.get("/", (req, res) => {
    res.sendFile("index.html", { root: "public" });
});

app.get("/forum", (req, res) => {
    res.sendFile("forum.html", { root: "public" });
});

app.get("/success", (req, res) => {
    res.sendFile("success.html", { root: "public" });
});

app.get("/cancel", (req, res) => {
    res.sendFile("cancel.html", { root: "public" });
});

// Stripe
const stripeGateway = new Stripe(process.env.STRIPE_API);
const DOMAIN = process.env.DOMAIN;

app.post("/stripe-checkout", async (req, res) => {
    const lineItems = req.body.items.map((item) => {
        const unitAmount = parseInt(item.price.replace(/[^0-9.-]+/g, '') * 100);
        return {
            price_data: {
                currency: 'usd',
                product_data: {
                    name: item.title,
                    images: [item.productImg]
                },
                unit_amount: unitAmount,
            },
            quantity: item.quantity,
        };
    });

    console.log('lineItems:', lineItems);

    // Checkout session
    const session = await stripeGateway.checkout.sessions.create({
        payment_method_types: ['card'],
        mode: 'payment',
        line_items: lineItems,
        success_url: `${DOMAIN}/success`,
        cancel_url: `${DOMAIN}/cancel`,
        billing_address_collection: 'required'
    });

    res.json({ url: session.url });
});

app.post('/submit-billing', (req, res) => {
    const { name, address, phone, items } = req.body;

    // Simpan data ke database
    const query = 'INSERT INTO billing_info (name, address, phone) VALUES (?, ?, ?)';
    db.query(query, [name, address, phone], (err, result) => {
        if (err) {
            console.error('Error saving billing info:', err);
            res.status(500).json({ success: false, error: 'Database error' });
        } else {
            res.json({ success: true });
        }
    });
});

app.post('/submit-forum', (req, res) => {
    const { question } = req.body;

    // Simpan data ke database
    const query = 'INSERT INTO forum (question) VALUES (?)';
    db.query(query, [question], (err, result) => {
        if (err) {
            console.error('Error saving forum info:', err);
            res.status(500).json({ success: false, error: 'Database error' });
        } else {
            res.json({ success: true });
        }
    });
});

app.listen(3000, () => {
    console.log("Server is running on port 3000");
});

axios.get('http://localhost:8080/data')
  .then(response => {
    console.log(response.data);
  })
  .catch(error => {
    console.error(`Error: ${error.message}`);
  });