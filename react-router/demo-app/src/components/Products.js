import {Link,Outlet} from 'react-router-dom'
function Products(){
    return (
        <>
        <div><h1>Products</h1></div>
        <nav>
        <Link to='featured'>Featured</Link>
        <Link to='general'>General</Link>
        </nav>
        <Outlet></Outlet>
        </>
    )
}
export default Products;