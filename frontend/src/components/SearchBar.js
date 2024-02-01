import style from '../pages/HomePage.module.css';
import searchImage from '../statics/search.png'

function SearchBar() {
  return (
    <>
	<div className={style.search}>
          <div className={style.searchbar}>
            <img src={searchImage} alt="serach" className={style.search_icon}/>
            <input type="text" placeholder="Search..." id="searchField"></input>
          </div>
    </div>
    </>
  )
}

export default SearchBar;
