import { FoodSetsInterface } from './IFoodSet';
import { BookingsInterface } from './IBooking';
export interface FoodOrderedsInterface {
   
    ID: number;
    Name: string;
    FoodTime: Date;
    TotalPrice: number;

    BookingID: number;
    Booking: BookingsInterface;

    FoodOrderedFoodSets: FoodOrderedFoodSetsInterface[];
}

export interface FoodOrderedFoodSetsInterface {
    ID: number;

    FoodSetID: number;
    FoodSet: FoodSetsInterface;

    FoodOrderedID: number;
    FoodOrdered: FoodOrderedsInterface;

    Quantity: number;
}