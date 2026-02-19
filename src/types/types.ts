export interface FormData {
  first_name: string
  last_name: string
  customer_number: string
  email: string
  post_code: string
  phone: number | null
  order_number: number | null
}

interface Address {
  post_code: string
  city: string
  street: string
}

export interface OrderItem {
  order_id: number
  product_id: number
  quantity: number
  price: number
  product: Product
}

export interface Order {
  ID: number,
  CreatedAt: string,
  customer_id: number
  order_number: number
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


export interface Product {
  product_uid: string
  name: string
  image: string
  product_type: string
  eans: string
  is_alcoholic: boolean
  UnitPrice: PriceInfo
  RetailPrice: PriceInfo
  categories: ProductCategory[]
}

interface PriceInfo {
  price: number
  measure: string
  measure_amount: number
}

interface ProductCategory {
  name: string
}