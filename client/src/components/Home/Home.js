import React, { Component } from "react";

import { map } from "underscore";
import { BrowserRouter, Link } from "react-router-dom";

import Header from "../Header/Header";

import { ReactCompomemt as RoomImg } from "./room.svg";

const SECTIONS = [
  { title: "Комната 1", href: "/room1", Icon: RoomImg },
  { title: "Комната 2", href: "/room2", Icon: RoomImg },
  { title: "Комната 3", href: "/room3", Icon: RoomImg },
  { title: "Комната 4", href: "/room4", Icon: RoomImg },
  { title: "Комната 5", href: "/room5", Icon: RoomImg },
  { title: "Комната 6", href: "/room6", Icon: RoomImg },
  { title: "Комната 7", href: "/room7", Icon: RoomImg },
  { title: "Комната 8", href: "/room8", Icon: RoomImg },
  { title: "Комната 9", href: "/room9", Icon: RoomImg },
  { title: "Комната 10", href: "/room10", Icon: RoomImg },
  { title: "Комната 11", href: "/room11", Icon: RoomImg },
  { title: "Комната 12", href: "/room12", Icon: RoomImg },
];

export default class Home extends Component {
  render() {
    return (
      <div className="Home">
        <Header />
        <div className="Home-Body">
          <div className="SectionNavigation">
            {map(SECTIONS, ({ title, href, Icon }) => (
              // с помощью компонента Link осуществляется
              // навигация по разделам приложения
              <BrowserRouter>
                <Link className="SectionNavigation-Item Section" to={href}>
                  <Icon className="Section-Icon" />
                  <span className="Section-Title">{title}</span>
                </Link>
              </BrowserRouter>
            ))}
          </div>
        </div>
      </div>
    );
  }
}
