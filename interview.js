// this has global scope
let carName = 'volvo'


// Function has access to carName
function myFunction() {
  return carName;
}

// recursion -> function that calls itself and has an ending condition

function findEnd(n) {
  if (n === 0) return
  console.log(n);


  return findEnd(n - 1)
}



// Part 1: JavaScript Fundamentals (ES6+)

//     Explain the difference between var, let, and const. When would you use one over the other?
// Var -> can be accessed in block scope
// let cant
// const cant any and also cant be re-assigned a value

//     What is the "Event Loop" in JavaScript? How does it handle asynchronous operations like setTimeout or API calls?
// The eventloop is the way js handled async code. When JS reaches an async function it stores it in the queue if the stack is not empty. While sync code is getting ran the async is running on side and once its done it gets put in the event queue. Once the stack is empty, the event loop puts it back into the stack and shows the result of the async function

//     Explain Closures. Can you give a practical example of when you might use one?
// Closures are functions that are defined inside another function. This is a way to encapsulate variables. It was the answer to "var" before let and const were introduced.

//     What is the difference between == and ===?
// == looks at value, === looks at also the type

//     How does the this keyword work? How does it behave differently in an Arrow function vs. a regular function?
//  In arrow functions the "this" keyword is bound to the lexical scope. In other words the surrounding block, while in a normal function "this" is defined by the function itself.

//     Destructuring & Spread: Given an object const user = { name: 'Alex', details: { age: 25 } }, how would you extract age into a variable? How would you create a shallow copy of this object while changing the name to 'Sam'?

//     Array Manipulation: I have an array of products. How would you filter out the ones that are out of stock, and then map them to return only their names? (Looking for chaining .filter().map()).

// Part 2: Vue.js (The Core)

//     Lifecycle Hooks:

//         What is the difference between created (or onCreated) and mounted (or onMounted)?

//         If you need to fetch data from an API, which hook would you usually use and why?

//     Reactivity: How does Vue track changes in data? (Bonus: Difference between Vue 2's Object.defineProperty and Vue 3's Proxy?)

//     Computed vs. Methods vs. Watchers:

//         When should you use a computed property instead of a method? (Keyword: Caching).

//         When is a watcher the right choice?

//     Component Communication:

//         Parent to Child: How do you pass data down? (Props)

//         Child to Parent: How do you communicate back up? (Emits/Events)

//         Grandparent to Grandchild: How would you handle this without passing props through every layer? (Provide/Inject or Store).

//     Directives: Explain how v-if differs from v-show. In what scenario is v-show better for performance?

//     Routing: How do you handle dynamic routes (e.g., /user/:id) in Vue Router? How do you access that ID inside the component?

// Part 3: TypeScript (Your specific focus)

//     Type vs. Interface:

//         Scenario: "I need to define the shape of a User object. Should I use type or interface?"

//         Follow-up: What can an interface do that a type cannot? (Answer: Declaration merging). What can a type do that interface cannot? (Answer: Unions, primitives, tuples).

//     Optional Properties: How do you mark a property in an interface as optional? (The ? syntax).

//     Union Types: How do you specify that a variable can be either a string or a number?

//     Any vs. Unknown: Why is using any generally discouraged? What is a safer alternative?

// Part 4: HTML5 & CSS3

//     Semantic HTML: Why is it better to use <button> or <a> tags instead of a <div> with an onClick event? (Accessibility/SEO).

//     The Box Model: If I have a div with width: 200px, padding: 20px, and border: 5px, what is the total width of the element? (Depends on box-sizing).

//     Flexbox vs. Grid:

//         How do you center an element vertically and horizontally using Flexbox?

//         When would you choose CSS Grid over Flexbox?

//     Specificity: If I have an ID selector #header and a class selector .header, which style wins? How do you calculate specificity?

// Part 5: General Web & Process

//     REST APIs:

//         What is the difference between POST and PUT?

//         What do status codes 200, 404, and 500 mean generally?

//     Git: You are working on a feature branch and the main branch has been updated. How do you bring those changes into your branch? (Merge vs Rebase).

//     Scenario: You find a bug in production. Walk me through your debugging process. (Reproduce -> Inspect -> Isolate -> Fix -> Test).
