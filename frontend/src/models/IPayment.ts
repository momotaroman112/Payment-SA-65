import { BillsInterface } from "./IBill";

export interface PaymentsInterface {
    ID: number;
    Name: string;

    Bill: BillsInterface[];
}