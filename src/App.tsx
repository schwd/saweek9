import React from "react";

import { BrowserRouter as Router, Switch, Route } from "react-router-dom";


import Navbar from "./components/Navbar";

import Users from "./components/Users";

import UserCreate from "./components/UserCreate";


export default function App() {

 return (

   <Router>

     <div>

       <Navbar />

       <Switch>

         <Route exact path="/" component={Users} />

         <Route exact path="/create" component={UserCreate} />

       </Switch>

     </div>

   </Router>

 );

}
