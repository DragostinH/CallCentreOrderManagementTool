interface Address {
  postCode: string
  city: string
  street: string
}
export interface Customer {
  id: number
  name: string
  address: Address
  customerNumber: string
  email: string
  telNumber: string
  orders?: []
}
