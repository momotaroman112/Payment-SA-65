import { BookingsInterface } from './IBooking';
import { EmployeesInterface } from './IEmployee';
import { FoodOrderedsInterface } from './IFoodordered';
import { PaymentsInterface } from './IPayment';
export interface BillsInterface {
    ID: number;
    BillTime: Date;
    EmployeeID: number;
    Employee: EmployeesInterface;

    PaymentTypeID: number;
    PaymentType: PaymentsInterface;

    BookingID: number;
    Booking : BookingsInterface;

    FoodOrderedID: number;
    FoodOrdered: FoodOrderedsInterface;

    TotalPrice: number;
}