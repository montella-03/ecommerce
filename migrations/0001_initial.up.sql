-- users
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       email TEXT UNIQUE NOT NULL,
                       password_hash TEXT NOT NULL,
                       created_at TIMESTAMP DEFAULT NOW()
);

-- products
CREATE TABLE products (
                          id SERIAL PRIMARY KEY,
                          name TEXT NOT NULL,
                          description TEXT,
                          price DECIMAL(10,2) NOT NULL,
                          stock INTEGER NOT NULL DEFAULT 0,
                          created_at TIMESTAMP DEFAULT NOW()
);

-- carts (simple – one active cart per user)
CREATE TABLE carts (
                       id SERIAL PRIMARY KEY,
                       user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
                       created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE cart_items (
                            id SERIAL PRIMARY KEY,
                            cart_id INTEGER REFERENCES carts(id) ON DELETE CASCADE,
                            product_id INTEGER REFERENCES products(id) ON DELETE CASCADE,
                            quantity INTEGER NOT NULL CHECK (quantity > 0),
                            UNIQUE(cart_id, product_id)
);

-- orders
CREATE TABLE orders (
                        id SERIAL PRIMARY KEY,
                        user_id INTEGER REFERENCES users(id),
                        total DECIMAL(10,2) NOT NULL,
                        status TEXT NOT NULL DEFAULT 'pending',   -- pending, paid, shipped, cancelled
                        created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE order_items (
                             id SERIAL PRIMARY KEY,
                             order_id INTEGER REFERENCES orders(id) ON DELETE CASCADE,
                             product_id INTEGER REFERENCES products(id),
                             quantity INTEGER NOT NULL,
                             price_at_time DECIMAL(10,2) NOT NULL
);