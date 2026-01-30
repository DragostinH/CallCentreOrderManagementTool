import { faker } from '@faker-js/faker';
import fs from 'fs';

// Generate customers
const customers = [];
for (let i = 1; i <= 50; i++) {
  customers.push({
    id: i,
    name: faker.person.fullName(),
    address: {
      postCode: faker.location.zipCode(),
      city: faker.location.city(),
      street: faker.location.streetAddress()
    },
    customerNumber: faker.string.alphanumeric(8).toUpperCase(),
    email: faker.internet.email(),
    telNumber: faker.phone.number(),
    orders: [] // will populate later
  });
}

// Generate products
const categories = [
  { id: 1, name: 'Beverages' },
  { id: 2, name: 'Food' },
  { id: 3, name: 'Household' },
  { id: 4, name: 'Electronics' },
  { id: 5, name: 'Clothing' }
];

const productTypes = ['BASIC'];
const measures = ['kg', 'l', 'each', 'pack'];

const products = [];
for (let i = 1; i <= 100; i++) {
  const unitPrice = faker.number.float({ min: 1, max: 100, precision: 0.01 });
  const measure = faker.helpers.arrayElement(measures);
  const measureAmount = measure === 'each' ? 1 : faker.number.int({ min: 1, max: 10 });
  const retailPrice = unitPrice * 1.2; // markup

  products.push({
    id: i,
    product_uid: faker.string.uuid(),
    favourite_uid: null,
    eans: [faker.string.numeric(13)],
    product_type: faker.helpers.arrayElement(productTypes),
    name: faker.commerce.productName(),
    image: faker.image.url(),
    unit_price: {
      price: unitPrice,
      measure: measure,
      measure_amount: measureAmount
    },
    retail_price: {
      price: retailPrice,
      measure: measure === 'each' ? 'each' : measure
    },
    is_alcoholic: faker.datatype.boolean(),
    reviews: {}, // empty for now
    categories: faker.helpers.arrayElements(categories, faker.number.int({ min: 1, max: 3 }))
  });
}

// Generate orders
const orders = [];
for (let i = 1; i <= 50; i++) {
  const customer = faker.helpers.arrayElement(customers);
  const numItems = faker.number.int({ min: 1, max: 5 });
  const items = [];
  let total = 0;
  for (let j = 0; j < numItems; j++) {
    const product = faker.helpers.arrayElement(products);
    const quantity = faker.number.int({ min: 1, max: 5 });
    const price = product.retail_price.price * quantity;
    total += price;
    items.push({
      productId: product.id,
      quantity: quantity,
      price: price
    });
  }
  orders.push({
    id: i,
    customerId: customer.id,
    orderDate: faker.date.past(),
    status: faker.helpers.arrayElement(['pending', 'shipped', 'delivered']),
    total: total,
    items: items
  });
  // Add order id to customer's orders
  customer.orders.push(i);
}

// Write to db.json
const db = {
  customers: customers,
  products: products,
  orders: orders,
  login: {} // keep existing
};

console.log(`Customers: ${customers.length}, Products: ${products.length}, Orders: ${orders.length}`);

fs.writeFileSync('db.json', JSON.stringify(db, null, 2));
console.log('Data generated and saved to db.json');
