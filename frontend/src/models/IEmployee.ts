import { BillsInterface } from "./IBill";

export interface EmployeesInterface {
    ID: number;
    Name: string;
    Email: string;
    Password: string;

    Bills: BillsInterface[];
}