interface Address {
  post_code: string
  city: string
  street: string
}

interface OrderItem {
  OrderID: number
  ProductID: number
  Quantity: number
  Price: number
}

export interface Order {
  CreatedAt: string,
  customer_id: number
  order_id: number
  order_date: string
  status: string
  total: number
  items: OrderItem[] | null
}
export interface Customer {
  ID: number
  first_name: string
  last_name: string
  Address: Address
  customer_number: string
  email: string
  phone: string
  orders: Order[]
}
