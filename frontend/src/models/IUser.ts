import { BillsInterface } from "./IBill";

export interface UserInterface {
    ID: number;
    Name: string;
    Email: string;
    Password: string;

    Bills: BillsInterface[];
}