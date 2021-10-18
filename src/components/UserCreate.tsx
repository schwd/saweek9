import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import {
  makeStyles,
  Theme,
  createStyles,
  alpha,
} from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import FormControl from "@material-ui/core/FormControl";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import Box from "@material-ui/core/Box";
import Typography from "@material-ui/core/Typography";
import Divider from "@material-ui/core/Divider";
import Snackbar from "@material-ui/core/Snackbar";
import Select from "@material-ui/core/Select";
import MuiAlert, { AlertProps } from "@material-ui/lab/Alert";


import React from 'react';
import TextField from '@material-ui/core/TextField';
import MenuItem from '@material-ui/core/MenuItem';
/* eslint-disable react/prop-types */
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import NoSsr from '@material-ui/core/NoSsr';


import { ReferInterface } from "../models/IRefer";
import { DoctorInterface } from "../models/IDoctor";
import { HospitalInterface } from "../models/IHospital";
import { MedicalHistoryInterface } from "../models/IMedicalHistory";
import { MedicalRecordInterface } from "../models/IMedicalRecord";

import {
  MuiPickersUtilsProvider,
  KeyboardDateTimePicker,
} from "@material-ui/pickers";
import DateFnsUtils from "@date-io/date-fns";
import Refers from "./Users";



const Alert = (props: AlertProps) => {
  return <MuiAlert elevation={6} variant="filled" {...props} />;
};

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1,
    },
    container: {
      marginTop: theme.spacing(2),
    },
    paper: {
      padding: theme.spacing(2),
      color: theme.palette.text.secondary,
    },
    text:{
        color: '#48D1CC',
        textAlign: 'center',
      },
      textdiag:{
        
      },
      combobox: {
        '& .MuiTextField-root': {
          margin: theme.spacing(2),
          width: '50ch',
        },
      },
      datetime: {

        margin: theme.spacing(2),
        width: 500,
      },
  
      textbox: {
        '& .MuiTextField-root': {
          margin: theme.spacing(2),
          width: '83ch',
        },
    },
  }),
);

function UserCreate() {
  const classes = useStyles();
  const [selectedDate, setSelectedDate] = useState<Date | null>(new Date());
  const [medicalrecords, setMedicalRecords] = useState<MedicalRecordInterface[]>([]);
  const [doctors, setDoctors] = useState<DoctorInterface[]>([]);
  const [hospitals, setHospitals] = useState<HospitalInterface[]>([]);
  const [medicalhistory, setMedicalHistory] = useState<MedicalHistoryInterface[]>([]);
  const [refer, setRefer] = useState<Partial<ReferInterface>>(
    {}
  );

  const [value, setValue] = React.useState('Controlled');
  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);

  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: { "Content-Type": "application/json" },
  };



  const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
    if (reason === "clickaway") {
      return;
    }
    setSuccess(false);
    setError(false);
  };

  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>
  ) => {
    const name = event.target.name as keyof typeof refer;
    setRefer({
      ...refer,
      [name]: event.target.value,
    });
  };

  const handleDateChange = (date: Date | null) => {
    console.log(date);
    setSelectedDate(date);
  };

const getDoctors = async () => {
    fetch(`${apiUrl}/doctors`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setDoctors(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getHospitals = async () => {
    fetch(`${apiUrl}/hospitals`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setHospitals(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getMedicalHistory = async () => {
    fetch(`${apiUrl}/medical_histories`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setMedicalHistory(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getMedicalRecords = async () => {
    fetch(`${apiUrl}/medical_records`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setMedicalRecords(res.data);
        } else {
          console.log("else");
        }
      });
  };


  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  const handleInputChange = (

    event: React.ChangeEvent<{ id?: string; value: any }>
 
  ) => {
 
    const id = event.target.id as keyof typeof UserCreate;
 
    const { value } = event.target;
 
    setRefer({ ...refer, [id]: value });
 
  };

function submit() {
  let data = {
    DoctorID: convertType(refer.DoctorID),
    MedicalRecordID: convertType(refer.MedicalRecordID),
    HospitalID: convertType(refer.HospitalID),
    MedicalHistoryID: convertType(refer.MedicalHistoryID),
    Cause: refer.Cause ?? "",
    Date: selectedDate,
  };

  const requestOptionsPost = {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  };

  fetch(`${apiUrl}/refers`, requestOptionsPost)
    .then((response) => response.json())
    .then((res) => {
      if (res.data) {
        setSuccess(true);
      } else {
        setError(true);
      }
    });
}

  
  useEffect(() => {
    getDoctors();
    getMedicalRecords();
    getHospitals();
    getMedicalHistory();
  }, []);



return (

    <div>
        <Snackbar open={success} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="success">
          บันทึกข้อมูลสำเร็จ
        </Alert>
      </Snackbar>
      <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="error">
          บันทึกข้อมูลไม่สำเร็จ
        </Alert>
      </Snackbar>
      <br/><br/> 
      <Typography variant="h5" className={classes.text}>
            แบบบันทึกการส่งต่อผู้ป่วยเกินศักยภาพ
          </Typography>
  
      <br/>
      <Grid className={classes.paper} container spacing={0}>
        
      <Grid item xs={3}></Grid>
        <Grid item xs={3}>
         แพทย์
         <form className={classes.combobox} noValidate >
          
         <Select
                
                native
                variant="outlined"
                value={refer.DoctorID}
                onChange={handleChange}
                inputProps={{
                  name: "DoctorID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกแพทย์
                </option>
                {doctors.map((item: DoctorInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>
    
         </form>
        </Grid>

        
        <Grid item xs={3} >
          ผู้ป่วย
          <form className={classes.combobox} noValidate>
          
          <Select
         
                native
                autoWidth
                variant="outlined"
                value={refer.MedicalRecordID}
                onChange={handleChange}
                inputProps={{
                  name: "MedicalRecordID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกผู้ป่วย
                </option>
                {medicalrecords.map((item: MedicalRecordInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.PatientName}
                  </option>
                ))}
              </Select>
   
        </form>
           
        
        </Grid>
        <Grid item xs={3}></Grid>
        
        <Grid item xs={3}></Grid>  
        <Grid item xs={3}>
         โรงพยาบาล
         <form className={classes.combobox} noValidate>
          
         <Select
         
         native
         variant="outlined"
         value={refer.HospitalID}
         onChange={handleChange}
         inputProps={{
           name: "HospitalID",
         }}
       >
         <option aria-label="None" value="">
           กรุณาเลือกโรงพยาบาล
         </option>
         {hospitals.map((item: HospitalInterface) => (
           <option value={item.ID} key={item.ID}>
             {item.Name}
           </option>
         ))}
       </Select>
    
         </form>
        </Grid>

        <Grid item xs={3}>
         ผลการวินิจฉัย
         <form className={classes.combobox} noValidate>
          
         <Select variant="outlined"                         
                            value={refer.MedicalHistoryID}
                            inputProps={{name: "MedicalHistoryID"}}
                            onChange={handleChange}
                            style={{ width: 400 }}
                        >
                            <MenuItem value={0} key={0}>กรุณาเลือกคำนำหน้า</MenuItem>
                            {medicalhistory.map((item: MedicalHistoryInterface) => (
                            <MenuItem value={item.ID} key={item.ID}>{item.Diagnosis}</MenuItem>))}
                        </Select>
    
         </form>
        </Grid>

        
        <Grid item xs={3}></Grid> 
        
        <Grid item xs={3}></Grid> 
        <Grid item xs={3} className={classes.textdiag}>
          สาเหตุ
        <div>
        <form className={classes.textbox} noValidate>
      <div>
        <TextField
          id="Cause"
          multiline
          rows={2}
          variant="outlined"
          value={refer.Cause || ""}
          onChange={handleInputChange}
        />
      </div>
    </form>
        </div>
        </Grid>
        <Grid item xs={6}></Grid>

        <Grid item xs={3}></Grid> 
    
        <Grid item xs={3}>
          วันที่
          <form className={classes.combobox} noValidate>
      

      <MuiPickersUtilsProvider utils={DateFnsUtils}>
                <KeyboardDateTimePicker
                  id="Date"
                  name="Date"
                  value={selectedDate}
                  onChange={handleDateChange}
                  label=""
                  minDate={new Date("2018-01-01T00:00")}
                  format="yyyy/MM/dd hh:mm a"
                />
              </MuiPickersUtilsProvider>
    </form>
        
        </Grid>
        <Grid item xs={7}>

        </Grid>
        <Grid item xs={3}>
        <Button/>
        </Grid>
      </Grid>
      <Grid item xs={12}>
            <Button
              component={RouterLink}
              to="/"
              variant="contained"
            >
              กลับ
            </Button>
            <Button
              style={{ float: "right" }}
              variant="contained"
              onClick={submit}
              color="primary"
            >
              บันทึก
            </Button>
          </Grid>
    </div>
  );
}

export default UserCreate;