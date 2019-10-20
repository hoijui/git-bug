import { makeStyles } from '@material-ui/styles';
import gql from 'graphql-tag';
import React from 'react';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';

const useStyles = makeStyles(theme => ({
  root: {
    display: 'flex',
    justifyContent: 'center',
  },
  container: {
    width: '100%',
    maxWidth: 800,
  },
  textField: {
    marginLeft: theme.spacing(1),
    marginRight: theme.spacing(1),
  },
  dense: {
    marginTop: theme.spacing(2),
  },
  button: {
    float: "right",
  }
}));

export default function NewBug() {
  const classes = useStyles();
  const [values, setValues] = React.useState({
    title: 'New Bug Title',
    description: '',
  });

  const handleChange = name => event => {
    setValues({ ...values, [name]: event.target.value });
  };

  return (
    <div className={classes.root}>
      <form className={classes.container} noValidate autoComplete="off">
        <div>
          <TextField
            id="title"
            label="Title"
            className={classes.textField}
            value={values.title}
            onChange={handleChange('title')}
            margin="normal"
            variant="outlined"
            fullWidth
          />
        </div>
        <div>
          <TextField
            id="description"
            label="Description"
            multiline
            rows="4"
            rowsMax="10"
            value={values.description}
            onChange={handleChange('description')}
            className={classes.textField}
            margin="normal"
            helperText="hello"
            variant="outlined"
            fullWidth
          />
        </div>
        <div>
          <TextField
            id="labels"
            label="Labels"
            type="search"
            className={classes.textField}
            margin="normal"
            variant="outlined"
            fullWidth
          />
        </div>
        <div>
          <Button variant="contained" className={classes.button}>
            Save Bug
          </Button>
        </div>
      </form>

    </div>
  );
}

