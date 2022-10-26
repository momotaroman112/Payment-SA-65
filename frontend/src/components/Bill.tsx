import React, { useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";

import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Paper from "@mui/material/Paper";
import Box from "@mui/material/Box";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";

import moment from "moment";
import { BillsInterface } from "../models/IBill";

function Bills() {
  const [bills, setBills] = React.useState<BillsInterface[]>([]);

  //เเก้เป็น getbill
  //รับข้อมูลมาจาก DB
  const getBills = async () => {
    const apiUrl = `http://localhost:8080/bills`;
    const requestOptions = {
      method: "GET",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
    };

    //การกระทำ
    fetch(apiUrl, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data);
        if (res.data) {
          setBills(res.data);
        } else {
          console.log("else");
        }
      });
  };

  //เพื่อให้มีการดึงข้อมูลใส่ combobox ตอนเริ่มต้นเเค่ครั้งเดียว
  //
  useEffect(() => {
    getBills();
  }, []);

  return (
    <div>
      <Container sx={{ marginTop: 2 }} maxWidth="lg">
        <Box display="flex">
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              ประวัติใบชำระเงิน
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/create"
              variant="contained"
              color="primary"
            >
              ใบชำระเงิน
            </Button>
          </Box>
        </Box>

        <TableContainer component={Paper} sx={{ marginTop: 4 }}>
          <Table sx={{ miinWidth: 650 }} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="left" width="15%">
                  ID
                </TableCell>
                <TableCell align="left" width="15%">
                  Employee Name
                </TableCell>
                <TableCell align="center" width="15%">
                  PaymentType
                </TableCell>
                <TableCell align="center" width="15%">
                  Booking
                </TableCell>
                {/* <TableCell align="center" width="20%">
                  Food Ordered
                </TableCell> */}
                <TableCell align="center" width="20%">
                  Total Price
                </TableCell>
                <TableCell align="center" width="20%">
                  Bill Time
                </TableCell>
              </TableRow>
            </TableHead>

            <TableBody>
              {bills.map((bill: BillsInterface) => (
                <TableRow key={bill.ID}>
                  <TableCell align="left">{bill.ID}</TableCell>
                  <TableCell align="left">{bill.Employee.Name}</TableCell>
                  <TableCell align="center">{bill.PaymentType.Name}</TableCell>
                  <TableCell align="center">{bill.Booking.Room}</TableCell>
                  {/* <TableCell align="center">{bill.FoodOrdered.Name}</TableCell> */}
                  <TableCell align="center">
                    {bill.TotalPrice}
                  </TableCell>
                  <TableCell align="center">
                    {moment(bill.BillTime).format("DD/MM/YYYY HH:mm")}
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Container>
    </div>
  );
}

export default Bills;
