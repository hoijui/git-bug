import React from 'react';
import Button from '@material-ui/core/Button';
import Menu from '@material-ui/core/Menu';
import MenuItem from '@material-ui/core/MenuItem';
import { makeStyles } from '@material-ui/styles';
import { Link } from 'react-router-dom';


const useStyles = makeStyles(theme => ({
  appTitle: {
    ...theme.typography.h6,
    color: 'white',
    textDecoration: 'none',
  },
}));

export default function SimpleMenu() {
  const classes = useStyles();
  const [anchorEl, setAnchorEl] = React.useState(null);

  const handleClick = event => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  return (
    <div>
      <Button aria-controls="simple-menu" aria-haspopup="true"
              onClick={handleClick} className={classes.appTitle}>
        ☰
      </Button>
      <Menu
        id="simple-menu"
        anchorEl={anchorEl}
        keepMounted
        open={Boolean(anchorEl)}
        onClose={handleClose}
      >
        <MenuItem onClick={handleClose}>
        <Link to="/newbug" className={classes.appTitle}>
            New Bug
        </Link>
        </MenuItem>
      </Menu>
    </div>
  );
}
