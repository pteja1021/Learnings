//import {Link} from 'react-router-dom'
import {NavLink} from 'react-router-dom'

function Navbar(){
    const navLinkStyles=({isActive})=>{ //isActive is passed by default from NavLink
        return {
            fontWeight:isActive?'bold':'italic',
            textDecoration:isActive?'none':'underline',
        }
    }
    return (
        <nav>
            <NavLink to='/' style={navLinkStyles}>Home </NavLink>
            <NavLink to='/about'style={navLinkStyles}>About </NavLink>
            <NavLink to='/products' style={navLinkStyles}>Products </NavLink>
        </nav>
    )
}
export default Navbar;