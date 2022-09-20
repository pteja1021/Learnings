import {Route,Routes} from 'react-router-dom'
import Home from './components/Home'
import About from './components/About'
import Products from './components/Products'
import Featured from './components/Featured'
import General from './components/General'
import Navbar from './components/Navbar'  //for the Links
import Notfound from './components/Notfound'

export default function App(){
  return (
    <> 
    <Navbar></Navbar>
    <Routes>
      <Route path='/' element={<Home />}></Route>
      <Route path='/about' element={<About />}></Route>
      <Route path='/products' element={<Products />}>
        <Route path='featured' element={<Featured/>}></Route>
        <Route path='general' element={<General/>}></Route>
      </Route>
      <Route path="*" element={<Notfound />}></Route>  {/* No route match */}
    </Routes>
    </>
  )
}