import IconButton from '@mui/material/IconButton';
import SearchIcon from "@mui/icons-material/Search";

const SearchButton = (props)=> (<IconButton
                    size="large"
                    edge="start"
                    color="inherit"
                    aria-label="search"
                    sx={{mr:2, ...props.style}}
                >
                <SearchIcon />
            </IconButton>)

export default SearchButton