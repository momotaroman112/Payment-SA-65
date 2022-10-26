import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from "react-router-dom";
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import IconButton from '@mui/material/IconButton';
import MenuIcon from '@mui/icons-material/Menu';
import MonetizationOnIcon from '@mui/icons-material/MonetizationOn';
import HistoryIcon from '@mui/icons-material/History';
import Drawer from "@mui/material/Drawer";
import List from "@mui/material/List"
import ListItem from "@mui/material/ListItem";
import ListItemIcon from "@mui/material/ListItemIcon";
import ListItemText from "@mui/material/ListItemText";
import HomeIcon from '@mui/icons-material/Home';
import { EmployeesInterface } from '../models/IEmployee';

export default function ButtonAppBar() {
  const [users, setUsers] = useState<EmployeesInterface>();
  const [openDrawer, setOpenDrawer] = useState(false);

  //ฟังชั่น เปิดปิด
  const toggleDrawer = (state: boolean) => (event: any) => {
    if (event.type === "keydown" && (event.key === "Tab" || event.key === "Shift")) {
      return;
    }
    setOpenDrawer(state);
  }

  const getEmployee = async() => {
    const id = localStorage.getItem("id")
    const apiUrl = `http://localhost:8080/employee/${id}`;    
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
        console.log(res.data);
        if(res.data) {
          const content = res.data;
          setUsers(content);
        } else {
          console.log("else");
        }
      });
  };
  
  const SignOut = () => {
    localStorage.clear();
    window.location.href = "/";
  }

  useEffect(() => {
    getEmployee();
  }, []);


  const menu = [
    { name: "หน้าแรก", icon: <HomeIcon  />, path: "/" },
    { name: "ใบชำระเงินของลูกค้า", icon: <MonetizationOnIcon  />, path: "/create" },
    { name: "ประวัติใบชำระเงิน", icon: <HistoryIcon  />, path: "/history" },
  ]
  return (
    <div style={{ flexGrow: 1 }}>
      <AppBar position="static">
        <Toolbar>
          <IconButton 
            edge="start" 
            sx={{ marginRight: 2 }}
            color="inherit" 
            aria-label="menu"
            onClick={toggleDrawer(true)} 
          >
            <MenuIcon />
          </IconButton>
          <Drawer open={openDrawer} onClose={toggleDrawer(false)}>
            <List 
              sx={{ width: 300 }}
              onClick={toggleDrawer(false)} 
              onKeyDown={toggleDrawer(false)}
            >
              {menu.map((item, index) => (
                <ListItem key={index} button component={RouterLink} to={item.path}>
                  <ListItemIcon>{item.icon}</ListItemIcon>
                  <ListItemText>{item.name}</ListItemText>
                </ListItem>
              ))}
            </List>
          </Drawer>
          <Typography variant="h6" sx={{ flexGrow: 1 }}>
            ระบบชำระเงินที่ใช้ในการจองห้อง
          </Typography>
          <div style={{marginRight: ".5rem"}}>
            <Typography align="right" variant="subtitle2">
              {users?.Email}
            </Typography>
          </div>
          <div>
            <Button onClick={SignOut}  color="inherit" style={{ marginRight: 12 }}>
              ออกจากระบบ
            </Button>
          </div>
        </Toolbar>
      </AppBar>
    </div>
  );
}
//<MenuItem onClick={SignOut}>Sign out</MenuItem>
//<Menu
// id="menu-appbar"
// anchorEl={anchorEl}
// anchorOrigin={{
//   vertical: 'top',
//   horizontal: 'right',
// }}
// keepMounted
// transformOrigin={{
//   vertical: 'top',
//   horizontal: 'right',
// }}
// open={Boolean(anchorEl)}
// onClose={handleClose}
// >
//</Menu>