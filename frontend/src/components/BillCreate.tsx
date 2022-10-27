import React, { useEffect } from "react";
import Box from "@mui/material/Box";
import Container from "@mui/material/Container";
import Divider from "@mui/material/Divider";
import FormControl from "@mui/material/FormControl";
import Grid from "@mui/material/Grid";
import MenuItem from "@mui/material/MenuItem";
import Paper from "@mui/material/Paper";
import Select, { SelectChangeEvent } from "@mui/material/Select";
import Snackbar from "@mui/material/Snackbar";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";

import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";

import { BillsInterface } from "../models/IBill";
import { BookingsInterface } from "../models/IBooking";
import { EmployeesInterface } from "../models/IEmployee";
import { FoodOrderedFoodSetsInterface, FoodOrderedsInterface } from "../models/IFoodordered";
import { PaymentsInterface } from "../models/IPayment";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";
import Button from "@mui/material/Button";
import MuiAlert, { AlertProps } from "@mui/material/Alert";

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
  props,
  ref
) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

export default function BillCreate() {
  const [employee, setEmployee] = React.useState<EmployeesInterface>();
  const [foodordered, setFoodordered] = React.useState<FoodOrderedsInterface[]>(
    []
  );
  const [booking, setBooking] = React.useState<BookingsInterface[]>([]);
  const [paymentTypes, setPaymentType] = React.useState<PaymentsInterface[]>(
    []
  );
  const [selectedBooking, setSelectedBooking] =
    React.useState<BookingsInterface>();

  const [bill, setBill] = React.useState<Partial<BillsInterface>>({
    BookingID: 0,
    FoodOrderedID: 0,
    PaymentTypeID: 0,
    TotalPrice: 0,
    BillTime: new Date(),
  });

  const [success, setSuccess] = React.useState(false);
  const [error, setError] = React.useState(false);

  const handleClose = (
    event?: Event | React.SyntheticEvent,
    reason?: string
  ) => {
    if (reason === "clickaway") {
      return;
    }
    setSuccess(false);
    setError(false);
  };

  const getEmployee = async () => {
    const apiUrl = `http://localhost:8080/employee/${localStorage.getItem(
      "id"
    )}`;

    const requestOptions = {
      method: "GET",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
    };

    fetch(apiUrl, requestOptions)
      .then((response) => response.json())

      .then((res) => {
        console.log(res);
        if (res.data) {
          setEmployee(res.data);
        } else {
          console.log(res.error);
        }
      });
  };

  const getFoodOrdered = async () => {
    const apiUrl = `http://localhost:8080/food_ordereds`;

    const requestOptions = {
      method: "GET",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
    };

    fetch(apiUrl, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res);
        if (res.data) {
          setFoodordered(res.data);
        } else {
          console.log(res.error);
        }
      });
  };
  const getBooking = async () => {
    const apiUrl = `http://localhost:8080/bookings`;

    const requestOptions = {
      method: "GET",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
    };

    fetch(apiUrl, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res);
        if (res.data) {
          setBooking(res.data);
        } else {
          console.log(res.error);
        }
      });
  };

  const getPaymentType = async () => {
    const apiUrl = `http://localhost:8080/payment_types`;

    const requestOptions = {
      method: "GET",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
    };

    fetch(apiUrl, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res);
        if (res.data) {
          setPaymentType(res.data);
        } else {
          console.log(res.error);
        }
      });
  };

  function submit() {
    let data = {
      EmployeeID: employee?.ID,
      BookingID: bill.BookingID,
      FoodOrderedID: booking.find((b) => b.ID === bill.BookingID)?.FoodOrdereds[0].ID,
      PaymentTypeID: bill.PaymentTypeID,
      BillTime: bill.BillTime,
      TotalPrice: bill.TotalPrice,
    };

    const apiUrl = `http://localhost:8080/bills`;
    const requestOptions = {
      method: "POST",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    };

    fetch(apiUrl, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res);
        if (res.data) {
          setSuccess(true);
        } else {
          setError(true);
          console.log(res.error);
        }
      });
  }

  const handleSelectChange = (event: SelectChangeEvent<number>) => {
    const name = event.target.name as keyof typeof bill;
    setBill({ ...bill, [name]: event.target.value });
  };

  const sumTotalPrice = () => {
    let bookingPrice =
      booking.find((b) => b.ID === bill.BookingID)?.Room.Type.Price ?? 0;
    let foodOrderedPrice =
      booking.find((b) => b.ID === bill.BookingID)?.FoodOrdereds[0].TotalPrice ?? 0;
      // foodordered.find((f) => f.ID === bill.FoodOrderedID)?.TotalPrice ?? 0;

    setBill({ ...bill, TotalPrice: bookingPrice + foodOrderedPrice });
  };

  useEffect(() => {
    getEmployee();
    getBooking();
    //getFoodOrdered();
    getPaymentType();
  }, []);

  useEffect(() => {
    sumTotalPrice();
  }, [bill.BookingID, bill.FoodOrderedID]);

  useEffect(() => {
    if (bill.BookingID === 0) return;
    setSelectedBooking(booking.find((b) => b.ID === bill.BookingID));
  }, [bill.BookingID]);

  return (
    <Container sx={{ marginTop: 2 }}>
      <Snackbar
        open={success}
        autoHideDuration={3000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="success">
          บันทึกข้อมูลสำเร็จ
        </Alert>
      </Snackbar>
      <Snackbar
        open={error}
        autoHideDuration={3000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="error">
          บันทึกข้อมูลไม่สำเร็จ
        </Alert>
      </Snackbar>
      <Paper>
        <Box
          display="flex"
          sx={{
            marginTop: 4,
          }}
        >
          <Box sx={{ paddingX: 2, paddingY: 2 }}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              รายการยืนยันการชำระเงินของสมาชิกระบบจองใช้ห้อง
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container sx={{ padding: 2 }}>
          <Grid item xs={4}>
            <p>พนักงานระบบ</p>
          </Grid>
          <Grid item xs={6}>
            <TextField disabled id="Name" value={employee?.Name} />
          </Grid>

          <Grid item xs={4}>
            <p>ห้องที่ต้องการชำระเงิน</p>
          </Grid>
          <Grid item xs={8}>
            <Select
              id="BookingID"
              value={bill.BookingID}
              inputProps={{
                name: "BookingID",
              }}
              onChange={handleSelectChange}
            >
              {booking.map((item: BookingsInterface) => (
                <MenuItem key={item.ID} value={item.ID}>
                  {item.Room.Name}
                </MenuItem>
              ))}
            </Select>
          </Grid>

          {/* <Grid item xs={4}>
            <p>ราคาห้อง</p>
          </Grid>
          <Grid item xs={6}>
            <TextField disabled  id = "Price" value={booking.map.room.Type.Price}/>
          </Grid> */}
          <Grid item xs={4}>
            <p>ชื่อผู้ใช้งาน</p>
          </Grid>
          <Grid item xs={8}>
            <TextField disabled id="Name" value={selectedBooking?.Member.Name} />
          </Grid>
          <Grid item xs={4}>
            <p>รายการสั่งอาหาร</p>
          </Grid>
          <Grid item xs={8}>
            {/* <Select
              id="FoodOrderedID"
              value={bill.FoodOrderedID}
              inputProps={{
                name: "FoodOrderedID",
              }}
              onChange={handleSelectChange}
            >
              {foodordered.map((item: FoodOrderedsInterface) => (
                <MenuItem key={item.ID} value={item.ID}>
                  {item.Name}
                </MenuItem>
              ))}
            </Select> */}
            <TableContainer component={Paper} sx={{ marginTop: 4 }}>
              <Table sx={{ miinWidth: 650 }} aria-label="simple table">
                <TableHead>
                  <TableRow>
                    <TableCell align="left" width="15%">
                      ID
                    </TableCell>
                    <TableCell align="center" width="20%">
                      Food Set Name
                    </TableCell>
                    <TableCell align="center" width="20%">
                      Quantity
                    </TableCell>
                    <TableCell align="center" width="20%">
                      Total Price
                    </TableCell>
                  </TableRow>
                </TableHead>

                <TableBody>
                  {selectedBooking?.FoodOrdereds[0].FoodOrderedFoodSets.map((foodOrder: FoodOrderedFoodSetsInterface) => (
                    <TableRow key={foodOrder.ID}>
                      <TableCell align="left">{foodOrder.FoodSet.ID}</TableCell>
                      <TableCell align="center">{foodOrder.FoodSet.Name}</TableCell>
                      <TableCell align="center">{foodOrder.Quantity}</TableCell>
                      <TableCell align="center">{foodOrder.Quantity * foodOrder.FoodSet.Price}</TableCell>
                      {/* 
                      <TableCell align="center">{foodOrder.TotalPrice}</TableCell>
                      <TableCell align="center">
                        {moment(foodOrder.FoodTime).format("DD/MM/YYYY HH:mm")}
                      </TableCell> */}
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            </TableContainer>
          </Grid>
          <Grid item xs={4}>
            <p>วิธีการชำระเงิน</p>
          </Grid>
          <Grid item xs={8}>
            <Select
              id="PaymentTypeID"
              value={bill.PaymentTypeID}
              inputProps={{
                name: "PaymentTypeID",
              }}
              onChange={handleSelectChange}
            >
              {paymentTypes.map((item: PaymentsInterface) => (
                <MenuItem key={item.ID} value={item.ID}>
                  {item.Name}
                </MenuItem>
              ))}
            </Select>
          </Grid>
          <Grid item xs={4}>
            <p>ราคารวม</p>
          </Grid>
          <Grid item xs={8}>
            <TextField disabled id="TotalPrice" value={bill?.TotalPrice} />
          </Grid>
          <Grid item xs={4}>
            <p>วันที่และเวลา</p>
          </Grid>
          <Grid item xs={8}>
            <FormControl fullWidth variant="outlined">
              <LocalizationProvider dateAdapter={AdapterDateFns}>
                <DatePicker
                  disabled
                  value={bill.BillTime}
                  onChange={(newValue) => {
                    setBill({
                      ...bill,
                      BillTime: newValue ? newValue : new Date(),
                    });
                  }}
                  renderInput={(params) => <TextField {...params} />}
                />
              </LocalizationProvider>
            </FormControl>
          </Grid>
          <Grid item xs={12}>
            <Button variant="contained" color="primary" onClick={submit}>
              ยืนยันการชำระเงิน
            </Button>
          </Grid>
        </Grid>
      </Paper>
    </Container>
  );
}
