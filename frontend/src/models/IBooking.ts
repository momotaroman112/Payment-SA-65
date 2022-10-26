import { FoodOrderedsInterface } from './IFoodordered';
import { UserInterface } from './IUser';
export interface BookingsInterface {
    ID: number;
    BookingTimeStart: Date;
    BookingTimeStop: Date;
    Room: string;
    TotalPrice: number;
    
    UserID: number;
    User: UserInterface;

    FoodOrdereds: FoodOrderedsInterface[];
    
    
}