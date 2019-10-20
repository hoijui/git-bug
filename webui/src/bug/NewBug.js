import { makeStyles } from '@material-ui/styles';
import gql from 'graphql-tag';
import React from 'react';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import { Mutation } from "react-apollo";

const QUERY = gql`
mutation($title: String!, $message: String!) {
  newBug(
   input: {
     title: $title,
     message: $message
 }) {
     bug {
       id,
       humanId
     },
    }
}
`;


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
  // react-apollo 3.1.2 I guess:
  // const [newBug, { data }] = useMutation(QUERY);
  const [values, setValues] = React.useState({
    title: 'New Bug Title',
    message: '',
  });

  const handleChange = name => event => {
    setValues({ ...values, [name]: event.target.value });
  };

  return (
    <Mutation mutation={QUERY}>
      {(newBug, { data }) => (
        <div className={classes.root}>
          <form className={classes.container} noValidate autoComplete="off"
            onSubmit={e => {
              e.preventDefault();
              newBug({
                variables: {
                  title: values.title,
                  message: values.message
                }
              });
            }}
          >
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
                id="message"
                label="message"
                multiline
                rows="4"
                rowsMax="10"
                value={values.message}
                onChange={handleChange('message')}
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
              <Button type="submit" variant="contained"
                className={classes.button}>
                Save Bug
          </Button>
            </div>
          </form>

        </div>
      )}
    </Mutation>
  );
}

